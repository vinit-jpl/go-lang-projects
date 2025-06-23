document.getElementById("uploadForm").addEventListener("submit", async function (e) {
    e.preventDefault();

    const fileInput = document.getElementById("fileInput");
    const formData = new FormData();
    formData.append("myFile", fileInput.files[0]);

    const response = await fetch("/upload", {
        method: "POST",
        body: formData,
    });

    const status = document.getElementById("status");
    if (response.ok) {
        const text = await response.text();
        status.innerText = `Success: \n ${text}`;
    } else {
        status.innerText = `Failed to upload file. The file should be of type image and should be less than 10MB.`;
    }
});
