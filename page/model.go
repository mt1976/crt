package page

import (
	lang "github.com/mt1976/crt/language"
	actn "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
)

// Page represents a page in a document or a user interface.
type Page struct {
	title            string         // The title of the page.
	pageRows         []pageRow      // The rows of content on the page.
	noRows           int            // The number of rows on the page.
	prompt           *lang.Text     // The prompt displayed to the user.
	showOptions      bool           // The text to be displayed to the user in the case options are possible
	actions          []*actn.Action // The available actions on the page.
	actionLen        int            // The maximum length of an action.
	blockedActions   []string       // The available actions on the page
	noPages          int            // The total number of pages.
	ActivePageIndex  int            // The index of the active page.
	counter          int            // A counter used for tracking.
	pageRowCounter   int            // A counter used for tracking the page rows.
	viewPort         *term.ViewPort // The viewPort object used for displaying the page.
	headerBarTop     int            // The header row top row
	headerBarContent int            // The header row content row
	headerBarBotton  int            // The header row bottom row
	footerBarTop     int            // The row where the input box starts
	footerBarInput   int            // The row where the input box is
	footerBarMessage int            // The row where the info box is
	footerBarBottom  int            // The last row of the page
	textAreaStart    int            // The row where the text area starts
	textAreaEnd      int            // The row where the text area ends
	height           int            // The height of the page
	width            int            // The width of the page
	maxContentRows   int            // The maximum number of rows available for content on the page.
	helpText         []string       // The help text to be displayed to the user
}

// pageRow represents a row of content on a page.
type pageRow struct {
	ID          int    // The unique identifier of the page row.
	RowContent  string // The content of the page row.
	PageIndex   int    // The index of the page row.
	Title       string // The title of the page row.
	AlternateID string // The alternate identifier of the page row.
	DateTime    string // The date and time of the page row.
}
