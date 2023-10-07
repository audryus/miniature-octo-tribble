'use strict';

const fs = require('fs');

let rawdata = fs.readFileSync('SI_C_OPCSEREMP.json');
let student = JSON.parse(rawdata);

const map1 = [];

function lero(item, idx) {
    let obj = map1.find(function(lItem, i){
        return lItem.razSocEmi.trim() === item.razSocEmi.trim() 
            && lItem.tipo.trim() === item.tipo.trim()
    })
    if (!obj) {
        map1.push({
            razSocEmi: item.razSocEmi.trim(),
            tipo: item.tipo.trim(),
        })
    }
}


student.Empresa["3"].forEach(lero) 
student.Empresa["A"].forEach(lero) 
student.Empresa["B"].forEach(lero) 
student.Empresa["C"].forEach(lero) 
student.Empresa["D"].forEach(lero) 
student.Empresa["E"].forEach(lero) 
student.Empresa["F"].forEach(lero) 
student.Empresa["G"].forEach(lero) 
student.Empresa["H"].forEach(lero) 
student.Empresa["I"].forEach(lero) 
student.Empresa["J"].forEach(lero) 
student.Empresa["K"].forEach(lero) 
student.Empresa["L"].forEach(lero) 
student.Empresa["M"].forEach(lero) 
student.Empresa["N"].forEach(lero) 
student.Empresa["O"].forEach(lero) 
student.Empresa["P"].forEach(lero) 
student.Empresa["Q"].forEach(lero) 
student.Empresa["R"].forEach(lero) 
student.Empresa["S"].forEach(lero) 
student.Empresa["T"].forEach(lero) 
student.Empresa["U"].forEach(lero) 
student.Empresa["V"].forEach(lero) 
student.Empresa["W"].forEach(lero) 
student.Empresa["X"].forEach(lero) 
student.Empresa["Y"].forEach(lero) 
student.Empresa["Z"].forEach(lero) 


fs.writeFileSync('empresa_tipo.json', JSON.stringify(map1));
