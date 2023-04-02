object HelloArgs {
  def main(args :Array[String]):Unit={
    if (args.length ==0){
      println("Hello World!")
    }else{
      val names :String = args.mkString(", ")
      println(s"Hello $names!")
    }
  }
}
