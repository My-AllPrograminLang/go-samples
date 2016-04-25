$(document).ready(function () {
 $("#output").append("Waiting for system time..");
 setInterval("delayedPost()", 1000);
});
function delayedPost() {
 $.post("http://localhost:9999/gettime", "", function(data, status) {
 $("#output").empty();
 $("#output").append(data);
 });
}
