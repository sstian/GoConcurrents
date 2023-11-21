/* 泛型：template 
C++ 的编译器会在编译时分析代码，根据不同的变量类型来自动化生成相关类型的函数或类，这叫模板的具体化。
*/

// 用<class T>来描述泛型
template <class T> 
T GetMax (T a, T b)  { 
  T result; 
  result = (a>b)? a : b; 
  return (result); 
} 

int main() {
  int i=5, j=6, k; 
  // 生成int类型的函数
  k=GetMax<int>(i,j);
  
  long l=10, m=5, n; 
  // 生成long类型的函数
  n=GetMax<long>(l,m); 

  return 0;
}

