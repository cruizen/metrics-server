// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package scraper

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper(in *jlexer.Lexer, out *Summary) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "node":
			(out.Node).UnmarshalEasyJSON(in)
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				in.Delim('[')
				if out.Pods == nil {
					if !in.IsDelim(']') {
						out.Pods = make([]PodStats, 0, 1)
					} else {
						out.Pods = []PodStats{}
					}
				} else {
					out.Pods = (out.Pods)[:0]
				}
				for !in.IsDelim(']') {
					var v1 PodStats
					(v1).UnmarshalEasyJSON(in)
					out.Pods = append(out.Pods, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper(out *jwriter.Writer, in Summary) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"node\":"
		out.RawString(prefix[1:])
		(in.Node).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"pods\":"
		out.RawString(prefix)
		if in.Pods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Pods {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Summary) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Summary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Summary) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Summary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper1(in *jlexer.Lexer, out *PodStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "podRef":
			(out.PodRef).UnmarshalEasyJSON(in)
		case "containers":
			if in.IsNull() {
				in.Skip()
				out.Containers = nil
			} else {
				in.Delim('[')
				if out.Containers == nil {
					if !in.IsDelim(']') {
						out.Containers = make([]ContainerStats, 0, 1)
					} else {
						out.Containers = []ContainerStats{}
					}
				} else {
					out.Containers = (out.Containers)[:0]
				}
				for !in.IsDelim(']') {
					var v4 ContainerStats
					(v4).UnmarshalEasyJSON(in)
					out.Containers = append(out.Containers, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper1(out *jwriter.Writer, in PodStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"podRef\":"
		out.RawString(prefix[1:])
		(in.PodRef).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"containers\":"
		out.RawString(prefix)
		if in.Containers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Containers {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper1(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper2(in *jlexer.Lexer, out *PodReference) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper2(out *jwriter.Writer, in PodReference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodReference) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodReference) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodReference) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodReference) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper2(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper3(in *jlexer.Lexer, out *NodeStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nodeName":
			out.NodeName = string(in.String())
		case "startTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "cpu":
			if in.IsNull() {
				in.Skip()
				out.CPU = nil
			} else {
				if out.CPU == nil {
					out.CPU = new(CPUStats)
				}
				(*out.CPU).UnmarshalEasyJSON(in)
			}
		case "memory":
			if in.IsNull() {
				in.Skip()
				out.Memory = nil
			} else {
				if out.Memory == nil {
					out.Memory = new(MemoryStats)
				}
				(*out.Memory).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper3(out *jwriter.Writer, in NodeStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix[1:])
		out.String(string(in.NodeName))
	}
	{
		const prefix string = ",\"startTime\":"
		out.RawString(prefix)
		out.Raw((in.StartTime).MarshalJSON())
	}
	if in.CPU != nil {
		const prefix string = ",\"cpu\":"
		out.RawString(prefix)
		(*in.CPU).MarshalEasyJSON(out)
	}
	if in.Memory != nil {
		const prefix string = ",\"memory\":"
		out.RawString(prefix)
		(*in.Memory).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper3(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper4(in *jlexer.Lexer, out *MemoryStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		case "workingSetBytes":
			if in.IsNull() {
				in.Skip()
				out.WorkingSetBytes = nil
			} else {
				if out.WorkingSetBytes == nil {
					out.WorkingSetBytes = new(uint64)
				}
				*out.WorkingSetBytes = uint64(in.Uint64())
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper4(out *jwriter.Writer, in MemoryStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.Raw((in.Time).MarshalJSON())
	}
	if in.WorkingSetBytes != nil {
		const prefix string = ",\"workingSetBytes\":"
		out.RawString(prefix)
		out.Uint64(uint64(*in.WorkingSetBytes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MemoryStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MemoryStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MemoryStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MemoryStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper4(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper5(in *jlexer.Lexer, out *ContainerStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "startTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "cpu":
			if in.IsNull() {
				in.Skip()
				out.CPU = nil
			} else {
				if out.CPU == nil {
					out.CPU = new(CPUStats)
				}
				(*out.CPU).UnmarshalEasyJSON(in)
			}
		case "memory":
			if in.IsNull() {
				in.Skip()
				out.Memory = nil
			} else {
				if out.Memory == nil {
					out.Memory = new(MemoryStats)
				}
				(*out.Memory).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper5(out *jwriter.Writer, in ContainerStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"startTime\":"
		out.RawString(prefix)
		out.Raw((in.StartTime).MarshalJSON())
	}
	if in.CPU != nil {
		const prefix string = ",\"cpu\":"
		out.RawString(prefix)
		(*in.CPU).MarshalEasyJSON(out)
	}
	if in.Memory != nil {
		const prefix string = ",\"memory\":"
		out.RawString(prefix)
		(*in.Memory).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper5(l, v)
}
func easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper6(in *jlexer.Lexer, out *CPUStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		case "usageCoreNanoSeconds":
			if in.IsNull() {
				in.Skip()
				out.UsageCoreNanoSeconds = nil
			} else {
				if out.UsageCoreNanoSeconds == nil {
					out.UsageCoreNanoSeconds = new(uint64)
				}
				*out.UsageCoreNanoSeconds = uint64(in.Uint64())
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
func easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper6(out *jwriter.Writer, in CPUStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.Raw((in.Time).MarshalJSON())
	}
	if in.UsageCoreNanoSeconds != nil {
		const prefix string = ",\"usageCoreNanoSeconds\":"
		out.RawString(prefix)
		out.Uint64(uint64(*in.UsageCoreNanoSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CPUStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CPUStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeSigsK8sIoMetricsServerPkgScraper6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CPUStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CPUStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeSigsK8sIoMetricsServerPkgScraper6(l, v)
}
