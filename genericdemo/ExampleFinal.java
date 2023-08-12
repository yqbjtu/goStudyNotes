
public class ExampleFinal {
  public static void main(String[] args) {
    System.out.println("finally 中没有return，但是修改了返回的变量的值（从1变为2）");
    System.out.println(testFinallyNoReturn());

    System.out.println("finally 中有return， finally中return的2， 非finally中return是1");
    System.out.println(testFinallyReturnValue());
  }

  public static int testFinallyReturnValue() {
    try {
      return 1;
    } finally {
      return 2;
    }
  }

  public static int testFinallyNoReturn() {
    int result = 1;

    try {
      return result;
    } finally {
      result = 2;
    }
  }
}

//最终输出结果
//finally 中没有return，但是修改了返回的变量的值（从1变为2）
//1 
//finally 中有return， finally中return的2， 非finally中return是1
//2
