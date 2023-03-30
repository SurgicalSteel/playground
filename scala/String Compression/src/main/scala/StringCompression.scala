object StringCompression {

  def main(args: Array[String]) {
    var s: String = scala.io.StdIn.readLine()
    var cc: String = s.substring(0, 1)
    var vs: String = ""
    var res: String = ""
    var ctr: Int = 0

    for (i <- 0 until s.length) {
      vs = s.substring(i, i + 1)

      if (cc != vs) {
        if (ctr == 1) {
          res += s"$cc"
        } else {
          res += s"$cc$ctr"
        }
        ctr = 1
        cc = vs
      } else {
        ctr += 1
      }

      if ((i + 1) == s.length) {
        if (ctr == 1) {
          res += s"$cc"
        } else {
          res += s"$cc$ctr"
        }
      }
    }
    println(res)
  }
}