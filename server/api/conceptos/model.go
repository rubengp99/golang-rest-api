package conceptos

type concepto struct{
	id int
    empresaID int
    codigo string
    referencia string
    nombre string
    descripcion string
    talla string
    color string
    descuento float32
    serialEstatico float32
    serialDinamico float32
    existenciaMinima float32
    existenciaMaxima float32
    tiposConceptosID int
    ubicacionID int
    costo float32
    ultimoCosto float32
    costoMayor float32
    costoPromedio float32
    fechaAt string
    fechaIn string
    fechaUc string
    gruposID int
    subgruposID int
    presentacion int
    unidadesID int
    fechaHora string
    marcasID int
    estado bool
    pvp float32
    precioA float32
    precionB float32
    precioDolar float32
    utilidad float32
    utiliadA float32
    utilidadB float32
    utilidadC float32 
    utilidadDolar float32
    costoDolar float32
    precioVariante float32
    retiene bool
    farmPrincipioActivoID int
    imagen string
    costoAdicional float32
    costoAdicional2 float32
    cantEnsamblado int
    licor bool
    porcentaje float32
    visiblePv bool
    visibleWeb bool
    restAreasID int
    setcortesia bool
    exento bool
    merma bool
    exitenciaC int
    obviarAjuste bool
    iva bool
    existencias []deposito
}

type deposito struct{
	admDepositosID int
	nombre string
	existencia float32
}