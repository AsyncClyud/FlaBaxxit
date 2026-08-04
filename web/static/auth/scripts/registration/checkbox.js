"use strict";

function SetAgreeValue() {
  const agree_button = document.getElementById("agree_button")
  agree_button.setAttribute("value", "agree")
  agree_button.setAttribute("onclick", "UnsetAgreeValue()")
}

function UnsetAgreeValue() {
  const agree_button = document.getElementById("agree_button")
  agree_button.setAttribute("value", "not agree")
  agree_button.setAttribute("onclick", "SetAgreeValue()")
}
