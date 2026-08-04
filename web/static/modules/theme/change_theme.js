"use strict";

function setTheme() {
  const saved = localStorage.getItem('theme');
  if (!saved) {
    document.documentElement.removeAttribute("class")
    document.documentElement.setAttribute("class", "dark")
    localStorage.setItem('theme', "dark");
  }
  else {
    if (saved == "light") {
      document.documentElement.removeAttribute("class")
      document.documentElement.setAttribute("class", "dark")
      localStorage.setItem('theme', "dark");
    }
    if (saved == "dark") {
      document.documentElement.removeAttribute("class")
      document.documentElement.setAttribute("class", "light")
      localStorage.setItem('theme', "light");
    }
  }
}
