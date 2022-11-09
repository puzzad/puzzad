<style>
    section {
        display: flex;
        justify-content: center;
    }

    @keyframes spinner {
        0% {
            transform: translate3d(-50%, -50%, 0) rotate(0deg);
        }
        100% {
            transform: translate3d(-50%, -50%, 0) rotate(360deg);
        }
    }

    .spin::before {
        animation: 1.5s linear infinite spinner;
        animation-play-state: inherit;
        border: solid 5px var(--border);
        border-bottom-color: var(--links);
        border-radius: 50%;
        content: '';
        height: 40px;
        width: 40px;
        position: absolute;
        transform: translate3d(-50%, -50%, 0);
        will-change: transform;
    }
</style>
<section>
    <div class="spin"></div>
</section>