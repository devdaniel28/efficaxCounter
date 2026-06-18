const style = require('./LoginOption.module.css')

export default function LoginOption({name, label, type, maxText, placeH, formValue}) {
    const labelText = `${label}:`
    const nameForm = String(name)
    const inputType = String(type)
    const placeHolder = String(placeH)
    const textMax = Number(maxText)
    
    return (
        <span className={style.login_opt}>
            <label htmlFor={nameForm} >{labelText}</label>
            
            <input type={inputType} 
                id={nameForm} 
                name={nameForm} 
                maxLength={textMax} 
                placeholder={placeHolder}
                value={formValue}
                required 
            />
        </span>
    )
}