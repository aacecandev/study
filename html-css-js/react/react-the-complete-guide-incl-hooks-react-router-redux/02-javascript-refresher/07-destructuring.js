// destructuring: easily extract array elements or object properties and store them in variables

// array destructuring
const numbers = [1, 2, 3];
[num1, , num3] = numbers;
console.log(num1, num3);

// object destructuring
const {name} = {name: 'Max', age: 28};
console.log(name);