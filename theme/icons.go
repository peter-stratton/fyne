package theme

import "fyne.io/fyne"

// ThemedResource is a resource wrapper that will return an appropriate resource
// for the currently selected theme.
type ThemedResource struct {
	source *fyne.DynamicResource
}

// Name returns the underlying resource name (used for caching)
func (res *ThemedResource) Name() string {
	return res.source.Name()
}

// Content returns the underlying content of the correct resource for the current theme
func (res *ThemedResource) Content() []byte {
	return res.source.Content()
}

// TODO: unfuck all of this stuff
// NewThemedResource creates a resource that adapts to the current theme setting.
func NewThemedResource(res *fyne.StaticResource) *ThemedResource {
	dr := &fyne.DynamicResource{
		BaseResource: res,
	}
	return &ThemedResource{
		source: dr,
	}
}

var (
	cancel, confirm, delete, search, searchReplace                              *ThemedResource
	checked, unchecked, radioButton, radioButtonChecked                         *ThemedResource
	contentAdd, contentRemove, contentCut, contentCopy, contentPaste            *ThemedResource
	contentRedo, contentUndo, info, question, warning                           *ThemedResource
	documentCreate, documentPrint, documentSave                                 *ThemedResource
	mailAttachment, mailCompose, mailForward, mailReply, mailReplyAll, mailSend *ThemedResource
	arrowBack, arrowDown, arrowForward, arrowUp                                 *ThemedResource
	folder, folderNew, folderOpen, help, home                                   *ThemedResource
	viewFullScreen, viewRefresh, viewZoomFit, viewZoomIn, viewZoomOut           *ThemedResource
)

func init() {
	cancel = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: cancelIconRes,
		},
	}
	confirm = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: checkIconRes,
		},
	}
	delete = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: deleteIconRes,
		},
	}
	search = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: searchIconRes,
		},
	}
	searchReplace = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: searchreplaceIconRes,
		},
	}

	checked = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: checkboxIconRes,
		},
	}
	unchecked = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: checkboxblankIconRes,
		},
	}
	radioButton = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: radiobuttonIconRes,
		},
	}
	radioButtonChecked = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: radiobuttoncheckedIconRes,
		},
	}

	contentAdd = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentaddIconRes,
		},
	}
	contentRemove = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentremoveIconRes,
		},
	}
	contentCut = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentcutIconRes,
		},
	}
	contentCopy = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentcopyIconRes,
		},
	}
	contentPaste = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentpasteIconRes,
		},
	}
	contentRedo = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentredoIconRes,
		},
	}
	contentUndo = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: contentundoIconRes,
		},
	}

	documentCreate = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: documentcreateIconRes,
		},
	}
	documentPrint = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: documentprintIconRes,
		},
	}
	documentSave = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: documentsaveIconRes,
		},
	}

	info = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: infoIconRes,
		},
	}
	question = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: questionIconRes,
		},
	}
	warning = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: warningIconRes,
		},
	}

	mailAttachment = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailattachmentIconRes,
		},
	}
	mailCompose = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailcomposeIconRes,
		},
	}
	mailForward = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailforwardIconRes,
		},
	}
	mailReply = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailreplyIconRes,
		},
	}
	mailReplyAll = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailreplyallIconRes,
		},
	}
	mailSend = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: mailsendIconRes,
		},
	}

	arrowBack = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: arrowbackIconRes,
		},
	}
	arrowDown = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: arrowdownIconRes,
		},
	}
	arrowForward = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: arrowforwardIconRes,
		},
	}
	arrowUp = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: arrowupIconRes,
		},
	}

	folder = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: folderIconRes,
		},
	}
	folderNew = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: foldernewIconRes,
		},
	}
	folderOpen = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: folderopenIconRes,
		},
	}
	help = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: helpIconRes,
		},
	}
	home = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: homeIconRes,
		},
	}

	viewFullScreen = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: viewfullscreenIconRes,
		},
	}
	viewRefresh = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: viewrefreshIconRes,
		},
	}
	viewZoomFit = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: viewzoomfitIconRes,
		},
	}
	viewZoomIn = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: viewzoominIconRes,
		},
	}
	viewZoomOut = &ThemedResource{
		&fyne.DynamicResource{
			BaseResource: viewzoomoutIconRes,
		},
	}
}

// FyneLogo returns a resource containing the Fyne logo
func FyneLogo() fyne.Resource {
	return fynelogo
}

// CancelIcon returns a resource containing the standard cancel icon for the current theme
func CancelIcon() fyne.Resource {
	return cancel
}

// ConfirmIcon returns a resource containing the standard confirm icon for the current theme
func ConfirmIcon() fyne.Resource {
	return confirm
}

// DeleteIcon returns a resource containing the standard delete icon for the current theme
func DeleteIcon() fyne.Resource {
	return delete
}

// SearchIcon returns a resource containing the standard search icon for the current theme
func SearchIcon() fyne.Resource {
	return search
}

// SearchReplaceIcon returns a resource containing the standard search and replace icon for the current theme
func SearchReplaceIcon() fyne.Resource {
	return searchReplace
}

// CheckButtonIcon returns a resource containing the standard checkbox icon for the current theme
func CheckButtonIcon() fyne.Resource {
	return unchecked
}

// CheckButtonCheckedIcon returns a resource containing the standard checkbox checked icon for the current theme
func CheckButtonCheckedIcon() fyne.Resource {
	return checked
}

// RadioButtonIcon returns a resource containing the standard radio button icon for the current theme
func RadioButtonIcon() fyne.Resource {
	return radioButton
}

// RadioButtonCheckedIcon returns a resource containing the standard radio button checked icon for the current theme
func RadioButtonCheckedIcon() fyne.Resource {
	return radioButtonChecked
}

// ContentAddIcon returns a resource containing the standard content add icon for the current theme
func ContentAddIcon() fyne.Resource {
	return contentAdd
}

// ContentRemoveIcon returns a resource containing the standard content remove icon for the current theme
func ContentRemoveIcon() fyne.Resource {
	return contentRemove
}

// ContentClearIcon returns a resource containing the standard content clear icon for the current theme
func ContentClearIcon() fyne.Resource {
	return cancel
}

// ContentCutIcon returns a resource containing the standard content cut icon for the current theme
func ContentCutIcon() fyne.Resource {
	return contentCut
}

// ContentCopyIcon returns a resource containing the standard content copy icon for the current theme
func ContentCopyIcon() fyne.Resource {
	return contentCopy
}

// ContentPasteIcon returns a resource containing the standard content paste icon for the current theme
func ContentPasteIcon() fyne.Resource {
	return contentPaste
}

// ContentRedoIcon returns a resource containing the standard content redo icon for the current theme
func ContentRedoIcon() fyne.Resource {
	return contentRedo
}

// ContentUndoIcon returns a resource containing the standard content undo icon for the current theme
func ContentUndoIcon() fyne.Resource {
	return contentUndo
}

// DocumentCreateIcon returns a resource containing the standard document create icon for the current theme
func DocumentCreateIcon() fyne.Resource {
	return documentCreate
}

// DocumentPrintIcon returns a resource containing the standard document print icon for the current theme
func DocumentPrintIcon() fyne.Resource {
	return documentPrint
}

// DocumentSaveIcon returns a resource containing the standard document save icon for the current theme
func DocumentSaveIcon() fyne.Resource {
	return documentSave
}

// InfoIcon returns a resource containing the standard dialog info icon for the current theme
func InfoIcon() fyne.Resource {
	return info
}

// QuestionIcon returns a resource containing the standard dialog question icon for the current theme
func QuestionIcon() fyne.Resource {
	return question
}

// WarningIcon returns a resource containing the standard dialog warning icon for the current theme
func WarningIcon() fyne.Resource {
	return warning
}

// FolderIcon returns a resource containing the standard folder icon for the current theme
func FolderIcon() fyne.Resource {
	return folder
}

// FolderNewIcon returns a resource containing the standard folder creation icon for the current theme
func FolderNewIcon() fyne.Resource {
	return folderNew
}

// FolderOpenIcon returns a resource containing the standard folder open icon for the current theme
func FolderOpenIcon() fyne.Resource {
	return folderOpen
}

// HelpIcon returns a resource containing the standard help icon for the current theme
func HelpIcon() fyne.Resource {
	return help
}

// HomeIcon returns a resource containing the standard home folder icon for the current theme
func HomeIcon() fyne.Resource {
	return home
}

// MailAttachmentIcon returns a resource containing the standard mail attachment icon for the current theme
func MailAttachmentIcon() fyne.Resource {
	return mailAttachment
}

// MailComposeIcon returns a resource containing the standard mail compose icon for the current theme
func MailComposeIcon() fyne.Resource {
	return mailCompose
}

// MailForwardIcon returns a resource containing the standard mail forward icon for the current theme
func MailForwardIcon() fyne.Resource {
	return mailForward
}

// MailReplyIcon returns a resource containing the standard mail reply icon for the current theme
func MailReplyIcon() fyne.Resource {
	return mailReply
}

// MailReplyAllIcon returns a resource containing the standard mail reply all icon for the current theme
func MailReplyAllIcon() fyne.Resource {
	return mailReplyAll
}

// MailSendIcon returns a resource containing the standard mail send icon for the current theme
func MailSendIcon() fyne.Resource {
	return mailSend
}

// MoveDownIcon returns a resource containing the standard down arrow icon for the current theme
func MoveDownIcon() fyne.Resource {
	return arrowDown
}

// MoveUpIcon returns a resource containing the standard up arrow icon for the current theme
func MoveUpIcon() fyne.Resource {
	return arrowUp
}

// NavigateBackIcon returns a resource containing the standard backward navigation icon for the current theme
func NavigateBackIcon() fyne.Resource {
	return arrowBack
}

// NavigateNextIcon returns a resource containing the standard forward navigation icon for the current theme
func NavigateNextIcon() fyne.Resource {
	return arrowForward
}

// ViewFullScreenIcon returns a resource containing the standard fullscreen icon for the current theme
func ViewFullScreenIcon() fyne.Resource {
	return viewFullScreen
}

// ViewRestoreIcon returns a resource containing the standard exit fullscreen icon for the current theme
func ViewRestoreIcon() fyne.Resource {
	return viewZoomFit
}

// ViewRefreshIcon returns a resource containing the standard refresh icon for the current theme
func ViewRefreshIcon() fyne.Resource {
	return viewRefresh
}

// ZoomFitIcon returns a resource containing the standard zoom fit icon for the current theme
func ZoomFitIcon() fyne.Resource {
	return viewZoomFit
}

// ZoomInIcon returns a resource containing the standard zoom in icon for the current theme
func ZoomInIcon() fyne.Resource {
	return viewZoomIn
}

// ZoomOutIcon returns a resource containing the standard zoom out icon for the current theme
func ZoomOutIcon() fyne.Resource {
	return viewZoomOut
}
