package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/goferwplynie/goXP/config"
)

func BuildStyle(s config.StyleConfig) lipgloss.Style {
	style := lipgloss.NewStyle()
	var borderZeroValue config.BorderConfig

	if s.ForegroundColor != "" {
		style = style.Foreground(lipgloss.Color(s.ForegroundColor))
	}
	if s.BackgroundColor != "" {
		style = style.Background(lipgloss.Color(s.ForegroundColor))
	}

	if s.Border != borderZeroValue {
		switch s.Border.BorderType {
		case "Block":
			style = style.Border(lipgloss.BlockBorder())
		case "Double":
			style = style.Border(lipgloss.DoubleBorder())
		case "InnerHalf":
			style = style.Border(lipgloss.InnerHalfBlockBorder())
		case "OuterHalf":
			style = style.Border(lipgloss.OuterHalfBlockBorder())
		case "Rounded":
			style = style.Border(lipgloss.RoundedBorder())
		case "Thick":
			style = style.Border(lipgloss.ThickBorder())
		}

		if !s.Border.Top {
			style = style.BorderTop(false)
		}
		if !s.Border.Right {
			style = style.BorderRight(false)
		}
		if !s.Border.Bottom {
			style = style.BorderBottom(false)
		}
		if !s.Border.Left {
			style = style.BorderLeft(false)
		}
	}
	if len(s.Padding) > 0 {
		style = style.Padding(s.Padding...)
	}
	if len(s.Margin) > 0 {
		style = style.Padding(s.Margin...)
	}
	if s.Bold {
		style = style.Bold(true)
	}

	return style
}
