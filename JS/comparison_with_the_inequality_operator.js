class Comparison {
  constructor(data) {
    this.value = data;
  }

  printData() {
    if (this.equalStricInequality()) {
      return "Not Equal";
    }
    return "Equal";
  }

  equalStricInequality() {
    return this.value !== 17;
  }
}

let compar = new Comparison("bob");
console.log(compar.printData());
