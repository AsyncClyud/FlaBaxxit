export async function IsAuth() {
  const response = await fetch("/api/auth", {
    method: "GET",
    headers: { Accept: "application/json" }
  })
  if (response.ok) {
    const data = await response.json()
    if (data.authorized == true) {
      return true
    }
    else {
      return false
    }
  }
}

export default IsAuth()
