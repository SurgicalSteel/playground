object PrefixCompression {

  def main(args: Array[String]) {
    var a: String = scala.io.StdIn.readLine()
    var b: String = scala.io.StdIn.readLine()
    var prefix: String = (a zip b).takeWhile(x => x._1 == x._2).map(_._1).mkString
    println(s"${prefix.length} $prefix")
    var subA: String = a.substring(prefix.length, a.length)
    var subB: String = b.substring(prefix.length, b.length)
    println(s"${a.length - prefix.length} $subA")
    println(s"${b.length - prefix.length} $subB")
  }
}