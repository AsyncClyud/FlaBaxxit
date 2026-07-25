"use strict";

import {IsAuth} from "../is_auth/is_auth.js"

async function UserMenu() {
  if (await IsAuth() == true) {
    const header_element = document.getElementById("header")

    const usermenu_button = document.createElement("button")
    usermenu_button.setAttribute("id", "usermenu_button")
    usermenu_button.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
    usermenu_button.setAttribute("onclick", "ShowUserMenu()")
    usermenu_button.textContent = "Menu"

    header_element.appendChild(usermenu_button)
  }
  else {
    const header_element = document.getElementById("header")

    const login_element = document.createElement("a")
    login_element.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
    login_element.href = `/auth/login`
    login_element.textContent = `Login`

    const register_element = document.createElement("a")
    register_element.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
    register_element.href = `/auth/register`
    register_element.textContent = `Registration`

    header_element.appendChild(login_element)
    header_element.appendChild(register_element)
  }
}

export default UserMenu()
