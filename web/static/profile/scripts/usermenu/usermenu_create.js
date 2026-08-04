"use strict";

import {IsAuth} from "../is_auth/is_auth.js"

async function UserMenu() {
  if (await IsAuth() == true) {
    const header_element = document.getElementById("header")

    const usermenu_button = document.createElement("button")
    usermenu_button.setAttribute("id", "usermenu_button")
    usermenu_button.setAttribute("class", "w-10 h-10 bg-[url(/static/images/3.webp)] bg-center bg-cover rounded-full m-[7px]")
    usermenu_button.setAttribute("onclick", "ShowUserMenu()")

    header_element.appendChild(usermenu_button)
  }
  else {
    const header_element = document.getElementById("header")

    const login_element = document.createElement("a")
    login_element.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
    login_element.href = `/auth/login`
    login_element.textContent = `Log in`

    const register_element = document.createElement("a")
    register_element.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
    register_element.href = `/auth/register`
    register_element.textContent = `Sing up`

    header_element.appendChild(login_element)
    header_element.appendChild(register_element)
  }
}

export default UserMenu()
