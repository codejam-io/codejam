<script lang="ts">
    let src='/favicon.png';
    export let size = 'normal'; // default size
</script>


<style>
    :root {
        --scale-factor: 1; /* Default scale factor */
    }

    .banner-tecj {
        height: 10rem;
        margin-bottom: 2rem;
        transform: scale(var(--scale-factor));
        transform-origin: top;
    }
    .banner-top {
        font-size: 2em;
        font-family: 'Liu Jian Mao Cao', cursive;
        text-align: center;
    }
    .banner-bottom {
        margin-top: -2rem;
        font-size: 2.5em;
        text-wrap: none;
        text-align: center;
    }
    .banner-bottom {
        text-shadow: 
            0 0 0.05em pink, /* Horizontal offset, vertical offset, blur radius, color */
            0.05em 0 0.05em pink,
            -0.05em 0 0.05em pink,
            0 0.05em 0.05em rgb(161, 107, 222),
            0 -0.05em 0.05em pink,
            0.05em 0.05em 0.05em rgb(55, 149, 193),
            -0.05em -0.05em 0.05em pink,
            0.05em -0.05em 0.05em pink,
            -0.05em 0.05em 0.05em pink;
    }
    .banner-tecj.small {
        --scale-factor: 0.5; /* Scale factor for small size */
    }
    .banner-top.small {
        font-size: 1rem;
    }
</style>

<div class="banner-tecj {`banner-tecj ${size}`}">
    <div class="banner-top mb-3">
        TimeEnjoyed's
    </div>
    <div class="banner-bottom">
        &lt;codejam/&gt;
    </div>
</div>