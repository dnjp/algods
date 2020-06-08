function fac(n) {
  if (n == 1) {
    return n;
  }
  return n * fac(n - 1);
}

function fac2(n, result) {
  if (n == 1) {
    return result;
  }
  return fac2(n - 1, result * n);
}

console.log(fac(5));
console.log(fac2(5, 1));