const letters = "абвгдежзиклмнопрстуфхцчшщъыьэюя".split('');

const results = [];
for (let i = 0; i <= 10000; i++) {
    results.push(getRandString());
}

console.log(results.join('\n'));

function getRandString() {
    const length = Math.round(1 + Math.random() * 3);
    let str = '';
    for (let i = 0; i < length; i++) {
        str += letters[Math.round(Math.random() * (letters.length - 1))];
    }

    return str;
}