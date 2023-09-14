class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        foodItems = set()
        tableFoodItems = {}
        for i in range(len(orders)):
            tableNumber = orders[i][1]
            foodItem = orders[i][2]
            foodItems.add(foodItem)
            if tableNumber not in tableFoodItems:
                tableFoodItems[tableNumber] = {
					foodItem: 1
				}
            elif foodItem not in tableFoodItems[tableNumber]:
                tableFoodItems[tableNumber][foodItem] = 1
            else:
                tableFoodItems[tableNumber][foodItem] += 1
        contents = []
        title = ["Table"]
        title.extend(sorted(list(foodItems)))
        for tableNumber, foodItemCount in tableFoodItems.items():
            content = [tableNumber]
            for foodItem in title[1:]:
                content.append(str(foodItemCount[foodItem]) if foodItem in foodItemCount else "0")
            contents.append(content)
        contents = sorted(contents, key=lambda x: int(x[0]))
        ans = []
        ans.append(title)
        ans.extend(contents)
        return ans