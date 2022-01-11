import React, { useEffect, useState } from 'react'
import { User } from '../entities/User'
import { ItemList } from '../components/ItemList'
import { WishlistItem } from '../entities/WistlistItem'

type Props = {
    user?: User
}

const Home: React.FC<Props> = ({user}) => {

    const [items, setItems] = useState<WishlistItem[]>([])

    useEffect(() => {
        if(user === undefined)
            return

        (async () => {
            const response = await fetch('/api/items');
            if(response.ok){
                const content = await response.json();
                setItems(content)
            }
        })()
    }, [user])

    const handleSubmitItem = async (item: WishlistItem) => {
        const response = await fetch('/api/items/', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(item)
        })

        if(response.ok){
            const content = await response.json()
            setItems([...items, content])
        }
    }

    if(user === undefined)
        return (<>please log in</>)

    return (
        <div>
            <h1>Hello, {user?.name}</h1>
            <ItemList items={items} onSubmitItem={handleSubmitItem} showPurchased={user.id === items[0]?.userId  } />
        </div>
    )
}

export default Home
