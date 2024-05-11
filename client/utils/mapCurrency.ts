import { OrderCurrency } from "../types/orderTypes";

export function mapCurrencySymbol(currency: OrderCurrency): string {
  switch (currency) {
    case OrderCurrency.UAH:
      return "₴";
    case OrderCurrency.USD:
      return "$";
    case OrderCurrency.EUR:
      return "€";
    default:
      throw new Error("Invalid currency");
  }
}
