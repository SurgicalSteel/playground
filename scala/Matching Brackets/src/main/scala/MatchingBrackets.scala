import scala.collection.mutable

object MatchingBrackets {
  def isPaired(brackets:String):Boolean = {
    var s = mutable.Stack[String]()
    var vs:String = ""
    var res:Boolean = true
    for(i <-0 until brackets.length){
      vs = brackets.substring(i,i+1)
      vs match {
        case "{" | "(" | "[" => s.push(vs)
        case "}" | ")" | "]" => if(s.size==0){res=false} else{res = isMatch(s.pop(), vs)}
        case _ =>
      }
      if (!res){
        res
      }
    }
    if(s.size >0){
      res =false
    }
    res
  }
  def isMatch(a:String,b:String):Boolean = (a=="(" && b== ")") || (a=="[" && b== "]") || (a=="{" && b== "}")

  def main(args:Array[String]):Unit={
    println(isPaired("([{}({}[])]){{}}"))
  }
}
