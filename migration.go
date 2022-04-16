package migrations

/**
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ``AS
 * IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL COPYRIGHT HOLDERS OR CONTRIBUTORS
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
**/

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"database/sql"
	"io/ioutil"
	"strconv"
	"strings"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type migration struct {
	path  string
	index int
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS.PRIVATE
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func newMigration(sqlPath string) (*migration, error) {
	var obj migration
	obj.path = sqlPath
	index, err := getMigrationIndexFromFile(sqlPath)
	if err != nil {
		return nil, err
	}
	obj.index = int(index)

	return &obj, nil
}

func (obj *migration) isExecuted(db *sql.DB) bool {
	row := db.QueryRow("SELECT * FROM `migrations` WHERE `index` = $1", obj.index)

	var index int
	err := row.Scan(&index)

	return err == nil
}

func (obj *migration) run(db *sql.DB) error {
	bytes, err := ioutil.ReadFile(obj.path)
	if err != nil {
		return err
	}

	query := string(bytes)
	_, err = db.Exec(query)

	return err
}

func getMigrationIndexFromFile(path string) (int, error) {
	nameParts := strings.Split(path, "_")
	nameParts = strings.Split(nameParts[0], "/")
	index, err := strconv.ParseInt(nameParts[len(nameParts)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	return int(index), err
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
