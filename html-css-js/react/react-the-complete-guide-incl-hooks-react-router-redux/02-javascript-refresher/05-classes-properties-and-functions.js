// ES6
// constructor() {
  // this.myProperty = 'value';
// }
// myMethod() {...}

// ES7
// myProperty = 'value';
// myMethod = () => {...}

class Human {
    gender = 'male';

  printGender = () => {
    console.log(this.gender);
  }
}

class Person extends Human {
    name = 'Max';
    gender = 'female';

  greet = () => {
    console.log('Hi, I am ' + this.name);
  }
}

const person = new Person();
person.greet();
person.printGender();