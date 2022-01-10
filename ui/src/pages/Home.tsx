import React, { useEffect, useState } from 'react'
import { User } from '../App'
import { ItemList } from '../components/ItemList'

type Props = {
    user?: User
}

const Home: React.FC<Props> = ({user}) => {

    const [loading, setLoading] = useState(true)

    useEffect(() => {
        if(user){
            setLoading(false)
        }
    }, [user])

    if(user == undefined)
        return (<>please log in</>)

    return (
        <div>
            <h1>Hello, {user?.name}</h1>
            <ItemList items={[]} />
        </div>
    )
}

export default Home
