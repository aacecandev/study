class Human {
  constructor() {
    this.gender = 'male';
  }

  printGender() {
    console.log(this.gender);
  }
}

class Person extends Human {
  constructor() {
    super();
    this.name = 'Max';
    this.gender = 'female';
  }

  greet() {
    console.log('Hi, I am ' + this.name);
  }
}

const person = new Person();
person.greet();
person.printGender();