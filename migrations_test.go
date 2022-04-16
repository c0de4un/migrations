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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type DBConfig struct {
	XMLName  xml.Name `xml:"DB"`
	Name     string   `xml:"name,attr"`
	Host     string   `xml:"host,attr"`
	Port     int      `xml:"port,attr"`
	Login    string   `xml:"login,attr"`
	Password string   `xml:"password,attr"`
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// UNITS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func TestMigrationRun(t *testing.T) {
	bytes, err := ioutil.ReadFile("configs/db.xml")
	if err != nil {
		t.Error(err)
	}

	// Read DB Config
	dbXML := &DBConfig{}
	err = xml.Unmarshal(bytes, dbXML)
	if err != nil {
		t.Error(err)
	}

	// Connect to DB
	db, err := loadDBConfig()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	testMigrateion, err := newMigration("db/migrations/0_up_init_migrations.sql")
	if err != nil {
		t.Error(err)
	}

	err = testMigrateion.run(db)
	if err != nil {
		t.Error(err)
	}
}

func TestUp(t *testing.T) {
	// Connect to DB
	db, err := loadDBConfig()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = Up("sql", 0, db)
	if err != nil {
		t.Error(err)
	}
}

func TestDown(t *testing.T) {
	// Connect to DB
	db, err := loadDBConfig()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = Down("sql", 0, db)
	if err != nil {
		t.Error(err)
	}
}

func loadDBConfig() (*sql.DB, error) {
	bytes, err := ioutil.ReadFile("configs/db.xml")
	if err != nil {
		return nil, err
	}

	// Read DB Config
	dbXML := &DBConfig{}
	err = xml.Unmarshal(bytes, dbXML)
	if err != nil {
		return nil, err
	}

	// Connect to DB
	src := fmt.Sprintf("%s:%s@/%s", dbXML.Login, dbXML.Password, dbXML.Name)
	db, err := sql.Open("mysql", src)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
