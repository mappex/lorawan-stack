// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

import fmt "fmt"

func (dst *Billing) SetFields(src *Billing, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "provider":
			if len(subs) == 0 && src == nil {
				dst.Provider = nil
				continue
			} else if len(subs) == 0 {
				dst.Provider = src.Provider
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "stripe":
					_, srcOk := src.Provider.(*Billing_Stripe_)
					if !srcOk && src.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'stripe', while different oneof is set in source")
					}
					_, dstOk := dst.Provider.(*Billing_Stripe_)
					if !dstOk && dst.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'stripe', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *Billing_Stripe
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Provider.(*Billing_Stripe_).Stripe
						}
						if dstOk {
							newDst = dst.Provider.(*Billing_Stripe_).Stripe
						} else {
							newDst = &Billing_Stripe{}
							dst.Provider = &Billing_Stripe_{Stripe: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Provider = src.Provider
						} else {
							dst.Provider = nil
						}
					}
				case "aws_saas_marketplace":
					_, srcOk := src.Provider.(*Billing_AWSSaaSMarketplace_)
					if !srcOk && src.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'aws_saas_marketplace', while different oneof is set in source")
					}
					_, dstOk := dst.Provider.(*Billing_AWSSaaSMarketplace_)
					if !dstOk && dst.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'aws_saas_marketplace', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *Billing_AWSSaaSMarketplace
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Provider.(*Billing_AWSSaaSMarketplace_).AWSSaaSMarketplace
						}
						if dstOk {
							newDst = dst.Provider.(*Billing_AWSSaaSMarketplace_).AWSSaaSMarketplace
						} else {
							newDst = &Billing_AWSSaaSMarketplace{}
							dst.Provider = &Billing_AWSSaaSMarketplace_{AWSSaaSMarketplace: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Provider = src.Provider
						} else {
							dst.Provider = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Billing_Stripe) SetFields(src *Billing_Stripe, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "customer_id":
			if len(subs) > 0 {
				return fmt.Errorf("'customer_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CustomerID = src.CustomerID
			} else {
				var zero string
				dst.CustomerID = zero
			}
		case "plan_id":
			if len(subs) > 0 {
				return fmt.Errorf("'plan_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PlanID = src.PlanID
			} else {
				var zero string
				dst.PlanID = zero
			}
		case "subscription_id":
			if len(subs) > 0 {
				return fmt.Errorf("'subscription_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SubscriptionID = src.SubscriptionID
			} else {
				var zero string
				dst.SubscriptionID = zero
			}
		case "subscription_item_id":
			if len(subs) > 0 {
				return fmt.Errorf("'subscription_item_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SubscriptionItemID = src.SubscriptionItemID
			} else {
				var zero string
				dst.SubscriptionItemID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Billing_AWSSaaSMarketplace) SetFields(src *Billing_AWSSaaSMarketplace, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "customer_identifier":
			if len(subs) > 0 {
				return fmt.Errorf("'customer_identifier' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CustomerIdentifier = src.CustomerIdentifier
			} else {
				var zero string
				dst.CustomerIdentifier = zero
			}
		case "product_code":
			if len(subs) > 0 {
				return fmt.Errorf("'product_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProductCode = src.ProductCode
			} else {
				var zero string
				dst.ProductCode = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
