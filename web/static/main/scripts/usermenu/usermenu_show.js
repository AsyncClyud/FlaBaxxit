"use strict";

async function ShowUserMenu() {
  const header_element = document.getElementById("header")

  const usermenu_div = document.createElement("div")
  usermenu_div.setAttribute("id", "usermenu")
  usermenu_div.setAttribute("class", "h-fit w-fit bg-[#333c46] absolute rounded-[5px] p-[10px]")

  usermenu_div.innerHTML = `
    <a class="bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]" href="/profile">Profile</a> <br>
    <a class="bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]" href="/article/create">Create Article</a> <br>
    <a class="bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]" href="/profile/settings">Settings</a> <br>
    <button class="bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]" type="button" onclick="Logout()">Logout</button>
    `
  const usermenu_button = document.getElementById("usermenu_button")
  usermenu_button.setAttribute("onclick", "HideUserMenu()")

  header_element.appendChild(usermenu_div)

}

async function HideUserMenu() {
  const usermenu_div = document.getElementById("usermenu")
  usermenu_div.remove()
  const usermenu_button = document.getElementById("usermenu_button")
  usermenu_button.setAttribute("onclick", "ShowUserMenu()")
}
