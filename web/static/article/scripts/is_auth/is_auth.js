export async function IsAuth() {
  const response = await fetch("/api/auth", {
    method: "GET",
    headers: { Accept: "application/json" }
  })
  if (response.ok) {
    const data = await response.json()
    if (data.authorized == true) {
      const comment_creator_element = document.getElementById("comment_creator")

      comment_creator_element.innerHTML = `
        <h2 class="text-[4vh] text-center font-[JetBrains_Mono] m-[10px]">Leave a comment</h2>
        <textarea class="w-[35vw] bg-[white] text-[black] resize-none rounded-[5px] m-[5px]" id="comment_content" rows="10" placeholder="Share your thoughts..."></textarea> <br>
        <button class="w-fit bg-[white] text-[black] text-[JetBrains_Mono] rounded-[3px] m-[5px] p-[5px]" type="button" onclick="SendCreateCommentRequest()">Post comment</button>
        <p id="status"></p>
        <div class="cf-turnstile" id="turnstile-widget" data-sitekey="0x4AAAAAAD9i83m-xS3y164A" data-theme="dark"></div>
        `
      return true
    }
    else {
      return false
    }
  }
}

export default IsAuth()
