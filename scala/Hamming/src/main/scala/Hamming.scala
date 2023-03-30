object Hamming {
  def distance(dnaStrandOne: String, dnaStrandTwo: String): Option[Int] = {
    if (dnaStrandOne.length != dnaStrandTwo.length){
      return None
    }
    var counter:Int = 0
    counter = dnaStrandOne zip dnaStrandTwo count(x => x._1 != x._2)
    Some(counter)
  }
  def main(args:Array[String]):Unit={
    println(distance("AAA","AAA").get)
    println(distance("AAA","AAC").get)
    println(distance("AAA","AAACT").getOrElse("None"))
  }
}
