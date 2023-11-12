// main.js
document.addEventListener("DOMContentLoaded", function () {
    loadFiles();
});

function loadFiles() {
    const fileSelect = document.getElementById("fileSelect");

    // Список файлов лежит локально, поэтому можем запросить его сразу
    fetchFilesList().then(data => {
        data.forEach(file => {
            const option = document.createElement("option");
            option.value = file;
            option.text = file;
            fileSelect.appendChild(option);
        });

        loadColorCodes();
    });
}

function fetchFilesList() {
files=this.files;
console.log("TEST");
 

const testFolder = './test/';
const fs = require('fs');

fs.readdirSync(testFolder).forEach(file => {
  console.log(file);
});

    const jsonFiles = [];

            for (let i = 0; i < files.length; i++) {
                console.log(files[i]);
                const file = files[i];
                if (file.type === 'application/json') {
                    jsonFiles.push(file.name);
                }
            }

    // Возвращаем Promise с массивом имен файлов
    return Promise.resolve(jsonFiles);
}

function loadColorCodes() {
    const colorCodeSelect = document.getElementById("colorCodeSelect");
    const selectedFile = document.getElementById("fileSelect").value;

    // Получаем список кодов краски для выбранного файла локально
    fetchColorCodes(selectedFile).then(data => {
        colorCodeSelect.innerHTML = "<option value='' disabled selected>Выберите код краски</option>";
        data.forEach(colorCode => {
            const option = document.createElement("option");
            option.value = colorCode;
            option.text = colorCode;
            colorCodeSelect.appendChild(option);
        });

        loadFormulas();
    });
}

function fetchColorCodes(file) {
    // Возвращаем Promise с массивом кодов краски
    // Здесь ты можешь использовать AJAX запрос или другой способ чтения файла
    // Для примера, возвращаем фиктивный список
    return Promise.resolve([
        "Code1",
        "Code2",
        "Code3",
        // Добавьте остальные коды краски
    ]);
}

// Аналогично адаптируйте loadFormulas и fetchFormulas функции
// ...

function displayFormulaDetails() {
    const selectedFile = document.getElementById("fileSelect").value;
    const selectedColorCode = document.getElementById("colorCodeSelect").value;
    const selectedFormula = document.getElementById("formulaSelect").value;

    // Получаем детали формулы локально
    fetchFormulaDetails(selectedFile, selectedColorCode, selectedFormula).then(data => {
        console.log("Formula details:", data);
    });
}

function fetchFormulaDetails(file, colorCode, formula) {
    // Возвращаем Promise с деталями формулы
    // Здесь также можешь использовать AJAX запрос или другой способ чтения файла
    // Для примера, возвращаем фиктивные данные
    return Promise.resolve({
        // Данные формулы
    });
}
