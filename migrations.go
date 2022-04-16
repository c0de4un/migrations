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
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func Up(dir string, version int, db *sql.DB) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), "_up_") {
			continue
		}

		migrationPath := dir + "/" + f.Name()
		fmt.Printf("\nRunning Up migration #%s\n", migrationPath)
		migration_, err := newMigration(migrationPath)
		if err != nil {
			return err
		}

		// Skip
		if (version > 0 && migration_.index < version) || migration_.isExecuted(db) {
			continue
		}

		// Run
		err = migration_.run(db)
		if err != nil {
			return err
		}

		// Save record about migration
		sqlStatement, err := db.Prepare("INSERT INTO `migrations` VALUES(?, ?)")
		if err != nil {
			return err
		}
		defer sqlStatement.Close()

		_, err = sqlStatement.Exec(migration_.index, f.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func Down(dir string, version int, db *sql.DB) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), "_down_") {
			continue
		}

		migrationPath := dir + "/" + f.Name()
		fmt.Printf("\nRunning Down migration #%s\n", migrationPath)
		migration_, err := newMigration(migrationPath)
		if err != nil {
			return err
		}

		// Skip
		if (version > 0 && migration_.index < version) || !migration_.isExecuted(db) {
			continue
		}

		// Run
		err = migration_.run(db)
		if err != nil {
			return err
		}

		// Save record about migration
		sqlStatement, err := db.Prepare("INSERT INTO `migrations` VALUES(?, ?)")
		if err != nil {
			return err
		}
		defer sqlStatement.Close()

		_, err = sqlStatement.Exec(migration_.index, f.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
