export const fetchEmpresasSemTipo = async () => {
    const empresas = await fetch("http://127.0.0.1:8080/api/v1/empresas")
    return await empresas.json()
}

export const patchEmpresasSemTipo = async (data) => {
    var json = JSON.stringify(data)
    const empresas = await fetch("http://127.0.0.1:8080/api/v1/empresas", {
        method: "PATCH", 
        body: json,
    })
    return await empresas.json()
}

export const fetchAtivos = async () => {
    const ativos = await fetch("http://127.0.0.1:8080/api/v1/ativos")
    return await ativos.json()
}