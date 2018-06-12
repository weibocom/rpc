// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	model "github.com/weibocom/ipc/model"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5a72dc82DecodeGithubComWeibocomIpcWebModel(in *jlexer.Lexer, out *Post) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.Post = new(model.Post)
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "similarity":
			out.Similarity = string(in.String())
		case "mid":
			out.MSGID = int64(in.Int64())
		case "dna":
			out.DNA = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "content":
			out.Content = string(in.String())
		case "content_type":
			out.ContentType = uint8(in.Uint8())
		case "keywords":
			out.Keywords = string(in.String())
		case "digest":
			out.Digest = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeGithubComWeibocomIpcWebModel(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"similarity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Similarity))
	}
	if in.MSGID != 0 {
		const prefix string = ",\"mid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MSGID))
	}
	if in.DNA != "" {
		const prefix string = ",\"dna\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DNA))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Content != "" {
		const prefix string = ",\"content\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Content))
	}
	if in.ContentType != 0 {
		const prefix string = ",\"content_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint8(uint8(in.ContentType))
	}
	if in.Keywords != "" {
		const prefix string = ",\"keywords\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Keywords))
	}
	if in.Digest != "" {
		const prefix string = ",\"digest\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Digest))
	}
	if true {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeGithubComWeibocomIpcWebModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeGithubComWeibocomIpcWebModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeGithubComWeibocomIpcWebModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeGithubComWeibocomIpcWebModel(l, v)
}
