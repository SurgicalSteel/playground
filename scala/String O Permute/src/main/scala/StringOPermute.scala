object StringOPermute {

  def process(s:String):String= s.grouped(2).toList.map(x => x.reverse).mkString

  def main(args:Array[String]):Unit={
    val nTestCases = scala.io.StdIn.readInt()
    var raw :String = ""
    for(i<-0 until nTestCases){
      raw = scala.io.StdIn.readLine()
      println(process(raw))
    }
    //println(process("abcdpqrs"))
  }

}
