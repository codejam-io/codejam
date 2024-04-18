<script>
    import LoginBtn from "./LoginBtn.svelte";
    let src = '../../static/pinktoasttransparent.png'
</script>
<style>
    .panel {
        height: 20vh;
        border: 0px solid blue;
        margin: 5rem 5rem;
        font-size: 1rem;
    }
    .panel-text {
        font-size: 2rem;
    }
    li {
        text-align: left;
        font-size: 2rem;
    }
    li::before {
        content: "✽";
        font-size: .8rem;
        color:rgb(171, 228, 255);
}
</style>

<div class="flex panel center">
    <div class="panel-text">
        Ingredients for Codejam:
        <ul>
            <li>
                sugar
            </li>
            <li>
                code
            </li>
        </ul>
        <div class="flex row sp-around">
            <LoginBtn />
        </div>
    </div>
    <img {src} alt="toast with jam and code on it">
</div>