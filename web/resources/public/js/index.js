document.body.addEventListener('htmx:beforeOnLoad', function (evt) {
    if (evt.detail.xhr.status === 400) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
    if (evt.detail.xhr.status === 401) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
    if (evt.detail.xhr.status === 500) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
});
htmx.onLoad(function(content) {
    let sortables = document.querySelectorAll(".sortable");
    for (let i = 0; i < sortables.length; i++) {
        let sortable = sortables[i];
        new Sortable(sortable, {
            animation: 150,
            ghostClass: 'blue-background-class'
        });
    }
});