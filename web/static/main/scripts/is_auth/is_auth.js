"use strict";

export async function IsAuth() {
  const response = await fetch("/api/auth", {
    method: "GET",
    headers: { Accept: "application/json" }
  })
  if (response.ok) {
    const data = await response.json()
    return data.userID
  }
  else {
    const data = await response.json()
    return data.userID
  }
}

export default IsAuth()
