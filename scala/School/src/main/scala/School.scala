object School {
  type DB = Map[Int, Seq[String]]
  var schoolDB: DB = Map()

  def add(name: String, g: Int): Unit = {
    var tempSeq = grade(g)
    tempSeq = tempSeq :+ name
    schoolDB = schoolDB updated(g, tempSeq)
  }

  def db: DB = schoolDB

  def grade(g: Int): Seq[String] = schoolDB.getOrElse(g, Seq.empty)

  def sorted: DB = {
    var keys = schoolDB.keys.toArray.sorted
    var tempMap: DB = Map()
    for (key <- keys) {
      var tempSeq = grade(key)
      tempSeq = tempSeq.sorted
      tempMap = tempMap updated(key, tempSeq)
    }
    tempMap
  }

  def main(args :Array[String]):Unit={
    add("Bob",1)
    add("Diana",1)
    add("Alice",2)
    add("Dave",2)
    add("Victor",3)
    println(grade(2))
    println(sorted)
  }

}