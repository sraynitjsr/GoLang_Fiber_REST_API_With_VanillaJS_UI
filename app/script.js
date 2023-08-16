$(document).ready(function() {
  const fetchButton = $("#fetch-button");

  fetchButton.click(function() {
    $.ajax({
      url: "http://localhost:3000",
      method: "GET",
      crossDomain: true,
      xhrFields: {
        withCredentials: true
      },
      headers: {
        "Content-Type": "application/json",
      },
      mode: "no-cors",
      success: function(data) {
        alert("Request successful, check browser console for more information.");
        console.log(data);
      },
      error: function(xhr, status, error) {
        console.error("Error fetching data:", error);
      }
    });
  });
});
