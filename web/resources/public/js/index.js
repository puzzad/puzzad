document.body.addEventListener('htmx:beforeOnLoad', function (evt) {
    if (evt.detail.xhr.status === 401) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
    if (evt.detail.xhr.status === 500) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
});