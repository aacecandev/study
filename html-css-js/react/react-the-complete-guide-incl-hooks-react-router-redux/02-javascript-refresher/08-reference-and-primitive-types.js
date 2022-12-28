// this creates a pointer to a memory address containing the person object
const person = {
  name: 'Max'
};

// this copies the above pointer to a new constant
const secondPerson = person;

console.log(secondPerson);
person.name = 'Manu';
console.log(secondPerson);

// this creates a pointer to a memory address containing the person object
const person2 = {
  name: 'Max'
};

// this copies the above memory address to a new constant
const secondPerson2 = {...person};

console.log(secondPerson2);
person.name = 'Manu';
console.log(secondPerson2);