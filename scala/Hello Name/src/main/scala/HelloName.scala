import scala.io.StdIn.readLine

object HelloName {
  def main(args : Array[String]) = {
    println("Enter your name :")
    var name :String = ""
    name = readLine()
    println(s"Hello, $name!")
  }
}
