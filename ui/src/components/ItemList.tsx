import { WishlistItem } from "../entities/WistlistItem"

type Props = {
    items: WishlistItem[]
}

export const ItemList: React.FC<Props> = ({items}) => {

    return (
        <div>
            Item List
        </div>
    )

}