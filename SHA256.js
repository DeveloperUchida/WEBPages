 // JavaScriptで動的にHTML要素を生成する
 window.onload = function() {
    var container = document.getElementById('container');
    var dynamicElement = document.createElement('p');
    dynamicElement.textContent = 'This paragraph is dynamically generated using JavaScript.';
    container.appendChild(dynamicElement);
  };