// Initialize EasyMDE markdown editor
const easyMDE = new EasyMDE({ element: document.getElementById("postContent") });

// Called when user clicks "Save Post"
function savePost() {
  const title = document.getElementById("title").value;
  const summary = document.getElementById("summary").value;
  const postMarkdown = easyMDE.value();

  console.log("Title:", title);
  console.log("Summary:", summary);
  console.log("Post Markdown:", postMarkdown);

  // Example: send to your Go backend
  fetch("/save-post", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      title: title,
      summary: summary,
      content: postMarkdown
    })
  })
  .then(res => res.json())
  .then(data => {
    alert("Post saved successfully!");
    console.log("Server response:", data);
  })
  .catch(err => console.error("Error saving post:", err));
}
