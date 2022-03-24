package main

type DB interface {
	Get() int
	Set() int
}

func TotalTime(db DB) int {
	return db.Get() - db.Set()
}

func main() {

}
