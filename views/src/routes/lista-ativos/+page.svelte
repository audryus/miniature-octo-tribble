<script>
    import { CodeBlock } from '@skeletonlabs/skeleton';

    /** @type {import('./$types').PageData} */
	export let data;
    const ativos = data.ativos
    let code = ""

    data.ativos.forEach((value, index) => {
        let symbol = value.symbol
        code += `${symbol}\t=RTD("rtdtrading.rtdserver";; "${symbol}_B_0"; "ULT")\n`
    })
    
</script>

<div class="container mx-auto p-8 space-y-8 ">
<CodeBlock code={code}>
</CodeBlock>

</div>