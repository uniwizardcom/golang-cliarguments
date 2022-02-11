package cliarguments

func (cli *LineService) convertLinkToItem(link LineServiceItemLink) *LineServiceItem {
	for _, item := range cli.itemsSupported {
		for _, linkRel := range item.Related {
			if link == linkRel {
				return &item
			}
		}
	}

	return nil
}

func (cli *LineService) convertItemToLink(item LineServiceItem) LineServiceItemLink {
	return LineServiceItemLink{
		Name:  item.Name,
		Level: item.Level,
	}
}
