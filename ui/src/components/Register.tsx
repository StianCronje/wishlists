import React, { SyntheticEvent, useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Register = () => {

    const [name, setName] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    const navigate = useNavigate();

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const request = {
            name,
            email,
            password
        }

        const response = await fetch('/api/register', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(request)
        })

        if(response.ok){
            navigate("/login")
        }
    }

    return (
        <div>
            <h2>Register</h2>
            <form onSubmit={submit}>
                <input type="text" name="name" placeholder='Name' value={name} onChange={(e) => setName(e.target.value)} />
                <input type="text" name="email" placeholder='Email' value={email} onChange={(e) => setEmail(e.target.value)} />
                <input type="password" name="password" placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)} />
                <button type="submit">Register</button>
            </form>            
        </div>
    )
}

export default Register
