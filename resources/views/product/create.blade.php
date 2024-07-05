<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <form action="{{ route('product.store') }}" method="post">
        @csrf
        Name<input type="text" name="product_name"><br>
        Price<input type="number" name="product_sell_price" step="0.01"><br>
        Measuring Unit<input type="text" name="measuring_unit"><br>
        MRP<input type="number" name="mrp" step="0.01"><br>
        category<input type="text" name="product_category"><br>
        quantity<input type="number" name="product_quantity"><br>
        Purchase Price<input type="number" name="purchase_price"><br>
        Ac sell Price<input type="number" name="ac_sale_price"><br>
        Non ac sell Price<input type="number" name="non_ac_sale_price"><br>
        Online delivery sale Price<input type="number" name="online_delivery_sell_price"><br>
        Online sale Price<input type="number" name="online_sell_price"><br>
        Tax<select name="tax" id="tax"><br>
        <option value="Non-GST-Supplies">Non-GST-Supplies</option>
        <option value="Exempted" name="Exempted">Exempted</option>
        <option value="GST@0%" name="GST@0%">GST@0%</option>
        <option value="GST@0.5%" name="GST@0.5%">GST@0.5%</option>
        <option value="GST@1%" name="GST@1%">GST@1%</option>
        <option value="GST@3%" name="GST@3%">GST@3%</option>
        <option value="GST@5%" name="GST@5%">GST@5%</option>
        <option value="GST@12%" name="GST@12%">GST@12%</option>
        <option value="GST@18%" name="GST@18%">GST@18%</option>
        <option value="GST@28%" name="GST@28%">GST@28%</option>
        </select>
        <br>Price with Tax
        <input type="radio" id="price_with_tax" name="price_with_tax" value="Y">Yes
        <input type="radio" id="price_with_tax" name="price_with_tax" value="N">No<br>
        Cess<input type="number" name="cess"><br>
        HSN code<input type="text" name="hsn_code"><br>
        Product description <input type="text" name="product_description"><br>
        Low Stock Alert <input type="number" name="low_stock_alert"><br>
        Product Storage Location <input type="text" name="product_storage_location"><br>
        Bulk purchase price <input type="number" name="bulk_purchase_price"><br>
        Retail sale unit price <input type="number" name="retail_sale_unit_per_bulk_purchase"><br>
        Bulk purchase unit price <input type="number" name="bulk_purchase_unit_per_retail_sale"><br>
        Expiry date <input type="date" name="expiry_date"><br>
        Show products online<input type="radio" id="show_product_online_store" name="show_product_online_store" value="Yes">Yes
        <input type="radio" id="show_product_online_store" name="show_product_online_store" value="No">No<br>
        <input type="submit">
    </form>
</body>
</html>