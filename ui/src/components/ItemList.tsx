import { SyntheticEvent, useState } from "react";
import { WishlistItem } from "../entities/WistlistItem";

type Props = {
  items: WishlistItem[]
  showPurchased: boolean
  onSubmitItem: (item: WishlistItem) => Promise<void> 
};

export const ItemList: React.FC<Props> = ({ items, showPurchased, onSubmitItem }) => {

    const [name, setName] = useState<string>('')
    const [price, setPrice] = useState<number>()
    const [link, setLink] = useState<string>('')

    const handleSubmit = (e: SyntheticEvent) => {
        e.preventDefault();

        const item: WishlistItem = {
            title: name,
            price: price ?? 0,
            link
        } as WishlistItem

        onSubmitItem(item)
    }

  return (
    <div>
      Item List
      <form onSubmit={handleSubmit}>
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Price</th>
              <th>Link</th>
              {showPurchased &&
                  <th>Purchased?</th>
              }
            </tr>
          </thead>
          <tbody>
            {items.map((item) => (
              <tr key={item.id}>
                <td>{item.title}</td>
                <td>{item.price}</td>
                <td>{item.link}</td>
                {showPurchased && 
                    <td>{item.isPurchased.toString()}</td>
                }
              </tr>
            ))}
            <tr>
              <td>
                <input placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
              </td>
              <td>
                <input type="number" placeholder="Price" value={price} onChange={(e) => setPrice(e.target.valueAsNumber)} />
              </td>
              <td>
                <input placeholder="Link" value={link} onChange={(e) => setLink(e.target.value)} />
              </td>
              <td>
                <button type="submit">Save</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  );
};
