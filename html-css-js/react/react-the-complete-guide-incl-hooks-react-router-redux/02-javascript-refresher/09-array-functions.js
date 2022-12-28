// map
const numbers = [1, 2, 3];
const doubleNumArray = numbers.map((num) => {
  return num * 2;
});
console.log(numbers);
console.log(doubleNumArray);

// filter
const filterArray = numbers.filter(num => {
  return num > 2;
});
console.log(filterArray);

// reduce
const reduceArray = numbers.reduce((prevValue, curValue) => {
  return prevValue + curValue;
});
console.log(reduceArray);