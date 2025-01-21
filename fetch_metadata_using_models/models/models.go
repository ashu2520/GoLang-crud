package models

type DataStoreSchema struct {
	DataTables      []*DataTables  `json:"dataTables,omitempty" yaml:"dataTables"`
	Constraints     []*Constraints `json:"constraints,omitempty" yaml:"constraints"`
	DataStore       string         `json:"dataStore,omitempty" yaml:"dataStore"`
	TrinoCatalog    string         `json:"vapusQueryServerUri,omitempty" yaml:"vapusQueryServerUri"`
	MetaSchemas     []string       `json:"metaSchemas,omitempty" yaml:"metaSchemas"`
	Description     string         `json:"description,omitempty" yaml:"description"`
	ComplianceTypes []string       `json:"complianceTypes,omitempty" yaml:"complianceTypes"`
}

func (x *DataStoreSchema) GetTableNames() []string {
	if x != nil {
		var res []string
		for _, val := range x.DataTables {
			res = append(res, val.Name)
		}
		return res
	}
	return []string{}
}

func (x *DataStoreSchema) GetFieldsList() []string {
	if x != nil {
		var res []string
		for _, val := range x.DataTables {
			for _, f := range val.Fields {
				res = append(res, f.Field)
			}
		}
		return res
	}
	return []string{}
}

func (x *DataStoreSchema) GetFedTableNames() map[string]string {
	if x != nil {
		res := map[string]string{}
		for _, val := range x.DataTables {
			res[val.FedTableName] = val.Name
		}
		return res
	}
	return map[string]string{}
}

type DataTables struct {
	Name                string        `json:"name,omitempty" yaml:"name"`
	Fields              []*DataFields `json:"fields,omitempty" yaml:"fields"`
	TotalRows           uint64        `json:"totalRows,omitempty" yaml:"totalRows"`
	TableType           string        `json:"tableType,omitempty" yaml:"tableType"`
	AverageRowLength    uint64        `json:"averageRowLength,omitempty" yaml:"averageRowLength"`
	IndexLength         uint64        `json:"indexLength,omitempty" yaml:"indexLength"`
	CreatedAt           uint64        `json:"createdAt,omitempty" yaml:"createdAt"`
	LastUpdatedAt       uint64        `json:"lastUpdatedAt,omitempty" yaml:"lastUpdatedAt"`
	DataLength          uint64        `json:"dataLength,omitempty" yaml:"dataLength"`
	Engine              string        `json:"engine,omitempty" yaml:"engine"`
	Version             string        `json:"version,omitempty" yaml:"version"`
	Nature              string        `json:"nature,omitempty" yaml:"nature"`
	TotalSize           uint64        `json:"totalSize,omitempty" yaml:"totalSize"`
	VapusQueryServerUri string        `json:"vapusQueryServerUri,omitempty" yaml:"vapusQueryServerUri"`
	GeneralUri          string        `json:"generalUri,omitempty" yaml:"generalUri"`
	Schema              string        `json:"schema,omitempty" yaml:"schema"`
	FedTableName        string        `json:"fedTableName,omitempty" yaml:"fedTableName"`
	Description         string        `json:"description,omitempty" yaml:"description"`
}

type DataFields struct {
	Name    string `json:"name,omitempty" yaml:"name"`
	Field   string `json:"field,omitempty" yaml:"field"`
	Type    string `json:"type,omitempty" yaml:"type"`
	Null    string `json:"null,omitempty" yaml:"null"`
	Key     string `json:"key,omitempty" yaml:"key"`
	Default string `json:"default,omitempty" yaml:"default"`
	Extra   string `json:"extra,omitempty" yaml:"extra"`
}

type Constraints struct {
	ConstraintType   string `json:"constraintType" yaml:"constraintType"`
	ConstraintName   string `json:"constraintName" yaml:"constraintName"`
	FieldName        string `json:"fieldName" yaml:"fieldName"`
	TableName        string `json:"tableName" yaml:"tableName"`
	TargetTable      string `json:"targetTable" yaml:"targetTable"`
	TargetColumn     string `json:"targetColumn" yaml:"targetColumn"`
	Enforced         bool   `json:"enforced" yaml:"enforced"`
	ConstraintSchema string `json:"constraintSchema" yaml:"constraintSchema"`
}
